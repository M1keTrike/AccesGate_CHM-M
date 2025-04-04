import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';
import { WebsocketService } from '../../services/websocket.service';

@Component({
    selector: 'app-scan-nfc',
    templateUrl: './scan-nfc.component.html',
    styleUrls: ['./scan-nfc.component.css'],
    standalone: false
})
export class ScanNfcComponent implements OnInit, OnDestroy {
  nfcCode: string = '';

  constructor(private router: Router, private wsService: WebsocketService) {}

  ngOnInit() {
    const deviceMac = localStorage.getItem('device_mac');
    if (!deviceMac) {
      console.error('❌ No se encontró la dirección MAC del dispositivo');
      return;
    }

    // Connect to WebSocket with NFC topic and device MAC
    this.wsService.connect('nfc', deviceMac);

    // Subscribe to messages
    this.wsService.getMessages().subscribe((uid: string) => {
      this.nfcCode = uid;
    });
  }

  ngOnDestroy() {
    // Cleanup WebSocket connection when component is destroyed
    this.wsService.disconnect();
  }

  confirmNFC() {
    localStorage.setItem('scannedNFC', this.nfcCode);
    this.router.navigate(['/admin/create-user']);
  }
}
