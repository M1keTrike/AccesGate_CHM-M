import { Component } from '@angular/core';
import { BluetoothService } from '../../../../services/Bluetooth.service';

@Component({
  selector: 'app-fingerprint-register',
  templateUrl: './fingerprint-register.component.html',
  styleUrls: ['./fingerprint-register.component.css'],
  standalone: false
})
export class FingerprintRegisterComponent {
  conectado = false;
  status: 'idle' | 'enviando' | 'exito' | 'error' = 'idle';
  mensaje: string = '';

  constructor(private bluetoothService: BluetoothService) {}

  async conectarDispositivo() {
    try {
      await this.bluetoothService.connect();
      this.conectado = true;
    } catch (err) {
      console.error('Error al conectar dispositivo BLE:', err);
      this.status = 'error';
      this.mensaje = '❌ Error al conectar con el dispositivo Bluetooth.';
    }
  }

  onStatusUpdate(event: { estado: string; mensaje: string }) {
    this.status = event.estado as any;
    this.mensaje = event.mensaje;
  }
}
