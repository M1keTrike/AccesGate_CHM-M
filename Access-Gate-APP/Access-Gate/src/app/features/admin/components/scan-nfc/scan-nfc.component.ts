import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WebsocketService } from '../../services/websocket.service';

@Component({
  selector: 'app-scan-nfc',
  templateUrl: './scan-nfc.component.html',
  styleUrls: ['./scan-nfc.component.css']
})
export class ScanNfcComponent implements OnInit {
  nfcCode: string = '';

  constructor(private router: Router, private wsService: WebsocketService) {}

  ngOnInit() {
    // Suscribirse al WebSocket y actualizar el campo NFC cuando llegue un mensaje
    this.wsService.getMessages().subscribe((uid: string) => {
      this.nfcCode = uid;
    });
  }

  confirmNFC() {
    localStorage.setItem('scannedNFC', this.nfcCode);
    this.router.navigate(['/admin/create-user']);
  }
}
