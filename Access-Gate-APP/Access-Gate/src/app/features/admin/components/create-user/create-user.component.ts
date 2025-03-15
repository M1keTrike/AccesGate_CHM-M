import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NfcCardsService } from '../../services/Nfc_cards.service';
import { UsersService } from '../../services/Users.Service';
import { NfcCard } from '../../models/iNfc_cards';
import { User } from '../../models/IUsers';

@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css']
})
export class CreateUserComponent implements OnInit {
  nfc: string = '';
  name: string = '';
  email: string = '';
  password: string = '';
  role: string = 'user';

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

  createUser() {
    const newUser: User = {
      id: 0, 
      name: this.name,
      email: this.email,
      password_hash: this.password, 
      role: this.role,
      created_at: new Date().toISOString()
    };
    console.log(newUser)
    return this.usersService.createUser(newUser);
  }

  onSubmit() {
    this.createNfcCard().subscribe(
      (nfcCard: NfcCard) => {
        console.log('NFC Card created:', nfcCard);
        this.createUser().subscribe(
          (user: User) => {
            console.log('User created:', user);
            this.router.navigate(['/']); // Navigate to the desired route after creation
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
