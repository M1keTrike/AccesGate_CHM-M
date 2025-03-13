import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css']
})
export class CreateUserComponent implements OnInit {
  nfc: string = '';

  constructor(private router: Router) {}

  ngOnInit() {
    const storedNFC = localStorage.getItem('scannedNFC');
    if (storedNFC) {
      this.nfc = storedNFC;
      localStorage.removeItem('scannedNFC'); // Limpiar el valor después de usarlo
    }
  }
  

  scanNFC() {
    this.router.navigate(['/admin/scan-nfc']);
  }
}
