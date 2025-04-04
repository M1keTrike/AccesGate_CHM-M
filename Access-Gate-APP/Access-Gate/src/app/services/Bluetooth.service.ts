import { Injectable } from '@angular/core';
import { DeviceService } from './device.service';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class BluetoothService {
  private device: BluetoothDevice | null = null;
  private server: BluetoothRemoteGATTServer | null = null;
  private service: BluetoothRemoteGATTService | null = null;

  // Store characteristics
  private ssidChar: BluetoothRemoteGATTCharacteristic | null = null;
  private passChar: BluetoothRemoteGATTCharacteristic | null = null;
  private statusChar: BluetoothRemoteGATTCharacteristic | null = null;
  private regFPChar: BluetoothRemoteGATTCharacteristic | null = null;
  private pirChar: BluetoothRemoteGATTCharacteristic | null = null;
  private ingresoChar: BluetoothRemoteGATTCharacteristic | null = null;

  // UUIDs remain the same
  private readonly WIFI_SERVICE_UUID = '12345678-1234-1234-1234-1234567890ab';
  private readonly WIFI_SSID_CHAR_UUID = '12345678-1234-1234-1234-1234567890ac';
  private readonly WIFI_PASS_CHAR_UUID = '12345678-1234-1234-1234-1234567890ad';
  private readonly WIFI_STATUS_CHAR_UUID = '12345678-1234-1234-1234-1234567890ae';
  private readonly REGISTER_FP_UUID = '12345678-1234-1234-1234-1234567890af';
  private readonly ENABLE_PIR_UUID = '12345678-1234-1234-1234-1234567890a2';
  private readonly ENABLE_INGRESO_UUID = '12345678-1234-1234-1234-1234567890a1';

  private statusCallback: ((value: string) => void) | null = null;

  async connect(): Promise<void> {
    try {
      this.device = await navigator.bluetooth.requestDevice({
        filters: [{ namePrefix: 'ESP32' }],
        optionalServices: [this.WIFI_SERVICE_UUID]
      });

      this.server = await this.device.gatt?.connect() || null;
      if (!this.server) throw new Error('Failed to connect to GATT server');

      this.service = await this.server.getPrimaryService(this.WIFI_SERVICE_UUID);
      if (!this.service) throw new Error('Service not found');

      // Get all characteristics in sequence
      try {
        this.ssidChar = await this.service.getCharacteristic(this.WIFI_SSID_CHAR_UUID);
        this.passChar = await this.service.getCharacteristic(this.WIFI_PASS_CHAR_UUID);
        this.statusChar = await this.service.getCharacteristic(this.WIFI_STATUS_CHAR_UUID);
        this.regFPChar = await this.service.getCharacteristic(this.REGISTER_FP_UUID);
        this.pirChar = await this.service.getCharacteristic(this.ENABLE_PIR_UUID);
        this.ingresoChar = await this.service.getCharacteristic(this.ENABLE_INGRESO_UUID);

        // Set up notifications using try-catch
        if (this.statusChar) {
          try {
            // @ts-ignore - Ignore TypeScript errors for Web Bluetooth API
            await this.statusChar.startNotifications();
            // @ts-ignore
            this.statusChar.addEventListener('characteristicvaluechanged', this.handleStatusChange.bind(this));
          } catch (notificationError) {
            console.warn('Notifications not supported:', notificationError);
          }
        }

        console.log('✅ Conectado al dispositivo:', this.device.name);
      } catch (error) {
        console.error('Error getting characteristics:', error);
        throw new Error('Failed to initialize BLE characteristics');
      }
    } catch (error) {
      console.error('❌ Error al conectar vía Bluetooth:', error);
      throw error;
    }
  }

  private handleStatusChange(event: any) {
    const value = new TextDecoder().decode(event.target.value);
    console.log('📡 Estado recibido:', value);
    if (this.statusCallback) {
      this.statusCallback(value);
    }
  }

  constructor(
    private deviceService: DeviceService,
    private authService: AuthService
  ) {}

  async sendWiFiCredentials(ssid: string, password: string): Promise<string | null> {
    if (!this.ssidChar || !this.passChar || !this.statusChar) {
      throw new Error('No hay conexión BLE o características no inicializadas');
    }

    return new Promise(async (resolve, reject) => {
      try {
        this.statusCallback = async (value: string) => {
          try {
            const data = JSON.parse(value);
            if (data.status === 'connected' && data.mac) {
              // Verify authentication token and get user ID
              const token = this.authService.getToken();
              if (!token) {
                reject(new Error('No hay sesión activa'));
                return;
              }

              const userId = this.authService.getUserId(); // Get user ID from decoded token
              if (!userId) {
                reject(new Error('ID de usuario no encontrado'));
                return;
              }

              // Verify device assignment
              try {
                await this.deviceService.verifyDeviceAssignment(data.mac, userId);
                const deviceData = {
                  hardware_id: data.mac,
                  assigned_to: userId,
                  type: 'ESP32',
                  status: 'active',
                  location: 'Default'
                };
                localStorage.setItem('pendingDevice', JSON.stringify(deviceData));
                resolve(data.mac);
              } catch (verifyError) {
                reject(verifyError);
              }
            } else {
              reject(new Error('No se pudo conectar al Wi-Fi'));
            }
          } catch (err) {
            reject(new Error('Respuesta inválida del ESP32'));
          }
          this.statusCallback = null;
        };

        await this.ssidChar!.writeValue(new TextEncoder().encode(ssid));
        await this.passChar!.writeValue(new TextEncoder().encode(password));

        setTimeout(() => {
          this.statusCallback = null;
          reject(new Error('Timeout esperando respuesta del ESP32'));
        }, 10000);

      } catch (err) {
        this.statusCallback = null;
        reject(err);
      }
    });
  }

  // Update other methods to use stored characteristics
  async setPir(enabled: boolean): Promise<void> {
    if (!this.pirChar) {
      throw new Error('No hay conexión BLE o característica no inicializada');
    }

    const value = new TextEncoder().encode(enabled ? '1' : '0');
    await this.pirChar.writeValue(value);
    console.log('PIR:', enabled ? 'habilitado' : 'deshabilitado');
  }

  async setAccessMode(enabled: boolean): Promise<void> {
    if (!this.ingresoChar) {
      throw new Error('No hay conexión BLE o característica no inicializada');
    }

    const value = new TextEncoder().encode(enabled ? '1' : '0');
    await this.ingresoChar.writeValue(value);
    console.log('Modo Ingreso:', enabled ? 'ACTIVO' : 'INACTIVO');
  }

  async sendRegisterFingerprintSignal(): Promise<void> {
    if (!this.regFPChar) {
      throw new Error('No hay conexión BLE o característica no inicializada');
    }

    const value = new TextEncoder().encode('1');
    await this.regFPChar.writeValue(value);
    console.log('Se ha solicitado el registro de huella.');
  }

  getDeviceName(): string | null {
    return this.device?.name || null;
  }

  isConnected(): boolean {
    return !!this.device && this.device.gatt?.connected === true;
  }
}
