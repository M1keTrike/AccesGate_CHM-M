import { Component } from '@angular/core';

@Component({
  selector: 'app-link-device',
  templateUrl: './link-device.component.html',
  styleUrls: ['./link-device.component.css'],
  standalone: false
})
export class LinkDeviceComponent {
  bluetoothStatus: string = 'disconnected';  
  macAddress?: string;

  onBluetoothConnected() {
    this.bluetoothStatus = 'connected';
  }

  onWiFiStatusUpdate(event: { status: string, mac?: string }) {
    this.bluetoothStatus = event.status;
    if (event.mac) {
      this.macAddress = event.mac;
    }
  }

  onWiFiConnected(mac: string) {
    this.bluetoothStatus = 'wifi-connected';
    this.macAddress = mac;
  }
  
}
