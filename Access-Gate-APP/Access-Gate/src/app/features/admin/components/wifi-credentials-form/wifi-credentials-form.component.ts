import { Component, EventEmitter, Output } from '@angular/core';
import { BluetoothService } from '../../../../services/Bluetooth.service';
import { DeviceService } from '../../../../services/device.service';

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

  constructor(
    private bluetoothService: BluetoothService,
    private deviceService: DeviceService
  ) {}

  async connectToWiFi() {
    this.isConnecting = true;
    this.errorMessage = '';

    try {
      const mac = await this.bluetoothService.sendWiFiCredentials(this.ssid, this.password);
      if (mac) {
        // Register the device with the API
        await this.deviceService.registerDevice(mac);
        this.wifiConnected.emit(mac);
      }
    } catch (error: any) {
      console.error('Error:', error);
      this.errorMessage = error.message;
    } finally {
      this.isConnecting = false;
    }
  }
}
