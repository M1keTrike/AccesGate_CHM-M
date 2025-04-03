import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SecurityService } from '../../services/security.service';
import { EventAttendee } from '../../models/event-attendee.model';

@Component({
  selector: 'app-event-attendees-control',
  templateUrl: './event-attendees-control.component.html',
  styleUrls: ['./event-attendees-control.component.css'],
  standalone: false
})
export class EventAttendeesControlComponent implements OnInit {
  eventId: number | null = null;
  attendees: EventAttendee[] = [];
  loading: boolean = true;

  constructor(
    private route: ActivatedRoute,
    private securityService: SecurityService
  ) {}

  ngOnInit(): void {
    this.eventId = Number(this.route.snapshot.paramMap.get('id'));
    if (this.eventId) {
      this.loadAttendees();
    }
  }

  loadAttendees(): void {
    if (!this.eventId) return;
    
    this.securityService.getEventAttendees(this.eventId).subscribe({
      next: (attendees) => {
        this.attendees = attendees;
        this.loading = false;
        console.log('Attendees loaded:', attendees);
      },
      error: (error) => {
        console.error('Error loading attendees:', error);
        this.loading = false;
      }
    });
  }

  updateAttendanceStatus(attendeeId: number, attended: boolean): void {
    if (!this.eventId) return;

    this.securityService.updateAttendanceStatus(this.eventId, attendeeId, attended).subscribe({
      next: () => {
        console.log('Attendance status updated successfully');
        this.loadAttendees(); // Reload the list after update
      },
      error: (error) => {
        console.error('Error updating attendance status:', error);
      }
    });
  }
}
