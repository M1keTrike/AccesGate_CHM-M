import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NfcCardsService } from '../../services/Nfc_cards.service';
import { UsersService } from '../../../../services/Users.Service';
import { NfcCard } from '../../models/iNfc_cards';
import { User } from '../../models/IUsers';

@Component({
    selector: 'app-create-user',
    templateUrl: './create-user.component.html',
    styleUrls: ['./create-user.component.css'],
    standalone: false
})
export class CreateUserComponent implements OnInit {
  nfc: string = '';
  name: string = '';
  email: string = '';
  password: string = '';
  role: string = '';

  constructor(
    private router: Router,
    private nfcCardsService: NfcCardsService,
    private usersService: UsersService
  ) {}

  ngOnInit() {
    this.getAllNfcCards();
    const storedNFC = localStorage.getItem('scannedNFC');
    if (storedNFC) {
      this.nfc = storedNFC;
      localStorage.removeItem('scannedNFC'); // Limpiar el valor después de usarlo
    }
  }
  
  getAllNfcCards() {
    this.nfcCardsService.getAllNfcCards().subscribe(
      (nfcCards: NfcCard[]) => {
        console.log('All NFC Cards:', nfcCards);
      },
      (error) => {
        console.error('Error fetching all NFC cards:', error);
      }
    );
  }

  scanNFC() {
    this.router.navigate(['/admin/scan-nfc']);
  }

  createNfcCard() {
    const newNfcCard: NfcCard = { card_uid: this.nfc };
    return this.nfcCardsService.createNfcCard(newNfcCard);
  }

  getCurrentUserId(): number {
    const token = localStorage.getItem('Authorization');
    if (token) {
      const tokenPayload = JSON.parse(atob(token.split('.')[1]));
      return tokenPayload.user_id;
    }
    throw new Error('No se encontró el token de autorización');
  }

  createUser() {
    try {
      const currentUserId = this.getCurrentUserId();
      const newUser: User = {
        id: 0,
        name: this.name,
        email: this.email,
        password_hash: this.password,
        role: this.role,
        created_at: new Date().toISOString(),
        created_by: currentUserId
      };
      console.log('Creando usuario:', newUser);
      return this.usersService.createUser(newUser);
    } catch (error) {
      console.error('Error al obtener el ID del usuario actual:', error);
      throw error;
    }
  }

  onSubmit() {
    this.createNfcCard().subscribe(
      (nfcCard: NfcCard) => {
        console.log('NFC Card created:', nfcCard);
        this.createUser().subscribe(
          (user: User) => {
            console.log('User created:', user);
            // Limpiar los campos después de crear el usuario exitosamente
            this.nfc = '';
            this.name = '';
            this.email = '';
            this.password = '';
            this.role = '';
          },
          (error) => {
            console.error('Error creating user:', error);
          }
        );
      },
      (error) => {
        console.error('Error creating NFC card:', error);
      }
    );
  }
}
