import { Component, EventEmitter, Output } from '@angular/core';
import { BluetoothService } from '../../../../services/Bluetooth.service';

@Component({
  selector: 'app-wifi-credentials-form',
  templateUrl: './wifi-credentials-form.component.html',
  styleUrls: ['./wifi-credentials-form.component.css'],
  standalone: false
})
export class WifiCredentialsFormComponent {
  ssid = '';
  password = '';
  isConnecting = false;
  errorMessage = '';

  @Output() wifiConnected = new EventEmitter<string>(); // MAC como string

  constructor(private bluetoothService: BluetoothService) {}

  async connectToWiFi() {
    this.isConnecting = true;
    this.errorMessage = '';

    try {
      const mac = await this.bluetoothService.sendWiFiCredentials(this.ssid, this.password);

      if (mac) {
        this.wifiConnected.emit(mac); // Emitir MAC al componente padre
      } else {
        this.errorMessage = 'Error: no se recibió dirección MAC del dispositivo.';
      }
    } catch (error) {
      console.error('Error al conectar al Wi-Fi:', error);
      this.errorMessage = 'No se pudo conectar a la red Wi-Fi.';
    } finally {
      this.isConnecting = false;
    }
  }
}
