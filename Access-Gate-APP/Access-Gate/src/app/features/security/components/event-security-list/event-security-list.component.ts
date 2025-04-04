import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SecurityService } from '../../services/security.service';
import { Event } from '../../models/event-attendee.model';
import { BluetoothService } from '../../../../services/Bluetooth.service';
import { EventAttendeeService } from '../../../admin/services/event-attendee.service';
import { EventService } from '../../../admin/services/event.service';
import { UsersService } from '../../../../services/Users.Service';

@Component({
  selector: 'app-event-security-list',
  templateUrl: './event-security-list.component.html',
  styleUrls: ['./event-security-list.component.css'],
  standalone: false
})
export class EventSecurityListComponent implements OnInit {
  events: Event[] = [];
  loading: boolean = true;
  isPirEnabled: boolean = false;
  isAccessModeEnabled: boolean = false;
  isConnected: boolean = false;
  isSecurityRole: boolean = true;
  idCreatedBy: number = 0;

  constructor(
    private router: Router,
    private securityService: SecurityService,
    private bluetoothService: BluetoothService,
    private eventService: EventService,
    private usersService: UsersService,
  ) {}

  ngOnInit(): void {
    this.loadEvents();
  }

  async connectBluetooth() {
    try {
      await this.bluetoothService.connect();
      this.isConnected = this.bluetoothService.isConnected();
    } catch (error) {
      console.error('Error connecting to Bluetooth:', error);
    }
  }

  async togglePir() {
    try {
      this.isPirEnabled = !this.isPirEnabled;
      await this.bluetoothService.setPir(this.isPirEnabled);
    } catch (error) {
      console.error('Error toggling PIR:', error);
      this.isPirEnabled = !this.isPirEnabled; // Revert state if failed
    }
  }

  async toggleAccessMode() {
    try {
      this.isAccessModeEnabled = !this.isAccessModeEnabled;
      await this.bluetoothService.setAccessMode(this.isAccessModeEnabled);
    } catch (error) {
      console.error('Error toggling access mode:', error);
      this.isAccessModeEnabled = !this.isAccessModeEnabled; // Revert state if failed
    }
  }

  private getCurrentUserId(): number {
    const token = localStorage.getItem('Authorization');
    if (token) {
      const tokenPayload = JSON.parse(atob(token.split('.')[1]));
      if(tokenPayload.role === 'security'){
        this.isSecurityRole=false
      }
      return tokenPayload.user_id;
    }
    return 0;
  }

  loadEvents(): void {
    this.loading = true;
    const currentUserId = this.getCurrentUserId();

    if (!this.isSecurityRole) {
      this.usersService.getUserById(currentUserId).subscribe({
        next: (user) => {
          this.idCreatedBy = user.created_by ?? 0;
          this.eventService.getAllEvents().subscribe({
            next: (events) => {
              this.events = events.filter(event => event.created_by === this.idCreatedBy);
              console.log('ID del creador del evento:', this.idCreatedBy);
              this.loading = false;
            },
            error: (error) => {
              console.error('Error loading events:', error);
              this.loading = false;
            }
          });
        },
        error: (error) => {
          console.error('Error loading user:', error);
        }
      });
    } else {
      this.eventService.getAllEvents().subscribe({
        next: (events) => {
          this.events = events.filter(event => event.created_by === currentUserId);
          console.log('ID del creador del evento:', currentUserId);
          this.loading = false;
        },
        error: (error) => {
          console.error('Error loading events:', error);
          this.loading = false;
        }
      });
    }
  }

  viewAttendees(eventId: number): void {
    this.router.navigate(['/organizer/security/event-attendees', eventId]);
  }
  logout() {
    this.usersService.logout();
    this.router.navigate(['/login']);
  }
}
