import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-link-status',
  templateUrl: './link-status.component.html',
  styleUrls: ['./link-status.component.css'],
  standalone: false
})
export class LinkStatusComponent {
  @Input() status: string = 'disconnected'; // 'disconnected' | 'connected' | 'wifi-connected'
  @Input() macAddress?: string;
}
