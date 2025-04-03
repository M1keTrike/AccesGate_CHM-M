import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SecurityService } from '../../services/security.service';
import { Event } from '../../models/event-attendee.model';
import { BluetoothService } from '../../../../services/Bluetooth.service';

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

  constructor(
    private router: Router,
    private securityService: SecurityService,
    private bluetoothService: BluetoothService
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

  loadEvents(): void {
    this.securityService.getAllEvents().subscribe({
      next: (events) => {
        this.events = events;
        this.loading = false;
      },
      error: (error) => {
        console.error('Error loading events:', error);
        this.loading = false;
      }
    });
  }

  viewAttendees(eventId: number): void {
    this.router.navigate(['/organizer/security/event-attendees', eventId]);
  }
}
