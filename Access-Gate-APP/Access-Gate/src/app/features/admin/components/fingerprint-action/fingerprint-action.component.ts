import { Component, EventEmitter, Output } from '@angular/core';
import { BluetoothService } from '../../../../services/Bluetooth.service';

@Component({
  selector: 'app-fingerprint-action',
  templateUrl: './fingerprint-action.component.html',
  styleUrls: ['./fingerprint-action.component.css'],
  standalone: false
})
export class FingerprintActionComponent {
  @Output() estado = new EventEmitter<{ estado: string, mensaje: string }>();

  constructor(private bluetoothService: BluetoothService) {}

  async iniciarRegistro() {
    this.estado.emit({ estado: 'enviando', mensaje: 'Enviando señal al dispositivo...' });

    try {
      await this.bluetoothService.sendRegisterFingerprintSignal();
      this.estado.emit({
        estado: 'exito',
        mensaje: '📲 Registro iniciado. Sigue las instrucciones en la pantalla del dispositivo.'
      });
    } catch (error) {
      console.error('❌ Error al iniciar registro:', error);
      this.estado.emit({
        estado: 'error',
        mensaje: '❌ No se pudo iniciar el registro de huella.'
      });
    }
  }
}
