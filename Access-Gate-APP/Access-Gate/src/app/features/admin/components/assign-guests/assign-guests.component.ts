import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { forkJoin } from 'rxjs';
import { MatSnackBar } from '@angular/material/snack-bar';
import { EventAttendeeService } from '../../services/event-attendee.service';
import { UsersService } from '../../../../services/Users.Service';
import { EventService } from '../../services/event.service';
import { Event } from '../../models/event';
import { User } from '../../models/IUsers';

@Component({
  selector: 'app-assign-guests',
  templateUrl: './assign-guests.component.html',
  styleUrls: ['./assign-guests.component.css'],
  standalone: false,
})
export class AssignGuestsComponent implements OnInit {
  assignForm: FormGroup;
  events: Event[] = [];
  users: User[] = [];
  attendees: User[] = [];
  filteredUsers: User[] = [];
  isLoading = false;
  selectedEventId: number | null = null;

  constructor(
    private fb: FormBuilder,
    private eventAttendeeService: EventAttendeeService,
    private userService: UsersService,
    private eventService: EventService,
    private snackBar: MatSnackBar
  ) {
    this.assignForm = this.fb.group({
      eventId: ['', Validators.required],
      userId: ['', Validators.required]
    });
  }

  token = localStorage.getItem('Authorization');

  ngOnInit(): void {
    if (!this.token) {
      this.showError('No authorization token found. Please login again.');
      return;
    }
    this.loadData();
  }

  private loadData(): void {
    this.isLoading = true;

    forkJoin({
      events: this.eventService.getAllEvents(),
      users: this.userService.getUsersByRole('attendee')
    }).subscribe({
      next: (result: { events: Event[], users: User[] }) => {
        this.events = result.events;
        this.users = result.users;
        this.filteredUsers = result.users;
        this.isLoading = false;
      },
      error: () => {
        this.showError('Error loading initial data');
        this.isLoading = false;
      }
    });
  }

  onEventSelect(event: EventTarget | any): void {
    const eventId = Number(event.target.value); // Extrae el ID correctamente
    if (!isNaN(eventId)) {
      this.selectedEventId = eventId;
      this.loadAttendees(eventId);
    }
  }

  private loadAttendees(eventId: number): void {
    this.isLoading = true;
    this.eventAttendeeService.getEventAttendees(eventId).subscribe({
      next: (attendees: User[]) => {
        this.attendees = attendees;
        this.filterAvailableUsers();
        this.isLoading = false;
      },
      error: () => {
        this.showError('Error loading attendees');
        this.isLoading = false;
      }
    });
  }

  private filterAvailableUsers(): void {
    const attendeeIds = this.attendees.map(a => a.id);
    this.filteredUsers = this.users.filter(user => !attendeeIds.includes(user.id));
  }

  assignGuest(): void {
    if (this.assignForm.valid) {
      this.isLoading = true;
      const { eventId, userId } = this.assignForm.value;

      // Ensure both IDs are numbers
      const newAttendee = {
        event_id: Number(eventId),
        user_id: Number(userId)
      };

      this.eventAttendeeService.registerAttendee(newAttendee).subscribe({
        next: (response) => {
          this.showSuccess('Guest assigned successfully');
          this.onEventSelect({ target: { value: eventId } } as unknown as Event);
          this.assignForm.get('userId')?.reset();  // Only reset the user selection
          this.isLoading = false;
        },
        error: (error) => {
          let errorMessage = 'Error assigning guest';
          if (error.error?.message) {
            errorMessage = error.error.message;
          } else if (error.status === 400) {
            errorMessage = 'Invalid data format or duplicate assignment';
          }
          this.showError(errorMessage);
          this.isLoading = false;
        }
      });
    }
  }

  removeAttendee(userId: number): void {
    if (this.selectedEventId) {
      this.isLoading = true;
      this.eventAttendeeService.removeAttendee(this.selectedEventId, userId).subscribe({
        next: () => {
          this.showSuccess('Attendee removed successfully');
          if (this.selectedEventId !== null) {
            this.loadAttendees(this.selectedEventId);
          }
        },
        error: () => {
          this.showError('Error removing attendee');
          this.isLoading = false;
        }
      });
    }
  }

  private showSuccess(message: string): void {
    this.snackBar.open(message, 'Close', { 
      duration: 3000,
      panelClass: ['success-snackbar']
    });
  }

  private showError(message: string): void {
    this.snackBar.open(message, 'Close', { 
      duration: 3000,
      panelClass: ['error-snackbar']
    });
  }
}