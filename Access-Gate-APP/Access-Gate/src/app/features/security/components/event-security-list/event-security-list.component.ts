import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SecurityService } from '../../services/security.service';
import { Event } from '../../models/event-attendee.model';

@Component({
  selector: 'app-event-security-list',
  templateUrl: './event-security-list.component.html',
  styleUrls: ['./event-security-list.component.css'],
  standalone: false
})
export class EventSecurityListComponent implements OnInit {
  events: Event[] = [];
  loading: boolean = true;

  constructor(
    private router: Router,
    private securityService: SecurityService
  ) {}

  ngOnInit(): void {
    this.loadEvents();
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
