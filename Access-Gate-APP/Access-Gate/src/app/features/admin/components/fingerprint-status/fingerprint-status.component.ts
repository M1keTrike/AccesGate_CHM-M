import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-fingerprint-status',
  templateUrl: './fingerprint-status.component.html',
  styleUrls: ['./fingerprint-status.component.css'],
  standalone: false
})
export class FingerprintStatusComponent {
  @Input() estado: string = 'idle';
  @Input() mensaje: string = '';
}
