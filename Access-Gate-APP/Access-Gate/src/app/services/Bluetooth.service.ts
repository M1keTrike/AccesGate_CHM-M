import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class BluetoothService {
  private device: BluetoothDevice | null = null;
  private server: BluetoothRemoteGATTServer | null = null;

  // UUID del servicio principal del ESP32
  private readonly WIFI_SERVICE_UUID = '12345678-1234-1234-1234-1234567890ab';

  // Características BLE (coinciden con las del ESP32)
  private readonly WIFI_SSID_CHAR_UUID = '12345678-1234-1234-1234-1234567890ac';
  private readonly WIFI_PASS_CHAR_UUID = '12345678-1234-1234-1234-1234567890ad';
  private readonly WIFI_STATUS_CHAR_UUID =
    '12345678-1234-1234-1234-1234567890ae';
  private readonly REGISTER_FP_UUID = '12345678-1234-1234-1234-1234567890af';

  // 1) UUID para habilitar/deshabilitar PIR
  private readonly ENABLE_PIR_UUID = '12345678-1234-1234-1234-1234567890ag';

  // 2) UUID para habilitar/deshabilitar "modo de ingreso"
  private readonly ENABLE_INGRESO_UUID = '12345678-1234-1234-1234-1234567890ah';

  async connect(): Promise<void> {
    if (!('bluetooth' in navigator)) {
      alert('Tu navegador no soporta Web Bluetooth. Usa Chrome o Edge.');
      return;
    }

    // Si ya está conectado, no volver a conectar
    if (this.device && this.device.gatt?.connected && this.server) {
      console.log('🔄 Ya hay una conexión activa con:', this.device.name);
      return;
    }

    try {
      this.device = await navigator.bluetooth.requestDevice({
        filters: [{ namePrefix: 'ESP32' }],
        // Aseguramos que se incluyan los servicios necesarios
        optionalServices: [this.WIFI_SERVICE_UUID],
      });

      this.server = (await this.device.gatt?.connect()) || null;
      console.log('✅ Conectado al dispositivo:', this.device?.name);
    } catch (error) {
      console.error('❌ Error al conectar vía Bluetooth:', error);
      throw error;
    }
  }

  async sendWiFiCredentials(
    ssid: string,
    password: string
  ): Promise<string | null> {
    if (!this.server) {
      throw new Error('No hay conexión Bluetooth');
    }

    try {
      const service = await this.server.getPrimaryService(
        this.WIFI_SERVICE_UUID
      );
      const ssidChar = await service.getCharacteristic(
        this.WIFI_SSID_CHAR_UUID
      );
      const passChar = await service.getCharacteristic(
        this.WIFI_PASS_CHAR_UUID
      );
      const statusChar = await service.getCharacteristic(
        this.WIFI_STATUS_CHAR_UUID
      );

      await ssidChar.writeValue(new TextEncoder().encode(ssid));
      await passChar.writeValue(new TextEncoder().encode(password));

      // Esperamos un breve lapso para que el ESP32 se conecte
      await new Promise((resolve) => setTimeout(resolve, 2500));

      const value = await statusChar.readValue();
      const decoded = new TextDecoder().decode(value).trim();

      if (!decoded) {
        throw new Error('No se recibió respuesta del ESP32');
      }

      let data: any;
      try {
        data = JSON.parse(decoded);
      } catch (err) {
        console.error('❌ JSON malformado recibido:', decoded);
        throw new Error('Respuesta inválida del ESP32');
      }

      if (data.status === 'connected') {
        return data.mac;
      } else {
        throw new Error('No se pudo conectar al Wi-Fi');
      }
    } catch (err) {
      console.error('❌ Error durante la transmisión de credenciales:', err);
      throw err;
    }
  }

  async sendRegisterFingerprintSignal(): Promise<void> {
    if (!this.server) {
      throw new Error('No hay conexión BLE');
    }

    const service = await this.server.getPrimaryService(this.WIFI_SERVICE_UUID);
    const characteristic = await service.getCharacteristic(
      this.REGISTER_FP_UUID
    );

    const value = new TextEncoder().encode('1');
    await characteristic.writeValue(value);

    console.log('Se ha solicitado el registro de huella.');
  }

  async setPir(enabled: boolean): Promise<void> {
    if (!this.server) {
      throw new Error('No hay conexión BLE');
    }

    // Tomamos el servicio principal
    const service = await this.server.getPrimaryService(this.WIFI_SERVICE_UUID);
    const pirCharacteristic = await service.getCharacteristic(
      this.ENABLE_PIR_UUID
    );

    // Convertimos '1' o '0' a bytes
    const value = new TextEncoder().encode(enabled ? '1' : '0');
    await pirCharacteristic.writeValue(value);

    console.log('PIR:', enabled ? 'habilitado' : 'deshabilitado');
  }

  async setAccessMode(enabled: boolean): Promise<void> {
    if (!this.server) {
      throw new Error('No hay conexión Bluetooth');
    }

    const service = await this.server.getPrimaryService(this.WIFI_SERVICE_UUID);
    const ingresoCharacteristic = await service.getCharacteristic(
      this.ENABLE_INGRESO_UUID
    );

    const value = new TextEncoder().encode(enabled ? '1' : '0');
    await ingresoCharacteristic.writeValue(value);

    console.log('Modo Ingreso:', enabled ? 'ACTIVO' : 'INACTIVO');
  }

  getDeviceName(): string | null {
    return this.device?.name || null;
  }

  isConnected(): boolean {
    return !!this.device && this.device.gatt?.connected === true;
  }
}
