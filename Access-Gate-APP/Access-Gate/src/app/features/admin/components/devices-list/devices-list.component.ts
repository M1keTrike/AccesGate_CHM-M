import { Component, EventEmitter, Output } from '@angular/core';
import { BluetoothService } from '../../../../services/Bluetooth.service';

@Component({
  selector: 'app-devices-list',
  templateUrl: './devices-list.component.html',
  styleUrls: ['./devices-list.component.css'],
  standalone: false
})
export class DevicesListComponent {
  @Output() deviceConnected = new EventEmitter<void>();

  isSearching = false;
  deviceName: string | null = null;

  constructor(private bluetoothService: BluetoothService) {}

  async scanAndConnect() {
    this.isSearching = true;
    try {
      await this.bluetoothService.connect();
      this.deviceName = this.bluetoothService.getDeviceName();
      this.deviceConnected.emit(); // Notifica al componente padre
    } catch (error) {
      console.error('Error al conectar con el dispositivo:', error);
    } finally {
      this.isSearching = false;
    }
  }
}
