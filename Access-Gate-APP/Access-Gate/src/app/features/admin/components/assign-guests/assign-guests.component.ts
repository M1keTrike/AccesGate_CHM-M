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
  token = localStorage.getItem('token');
  ngOnInit(): void {
    console.log(this.token)
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
      error: (err: any) => {
        this.showError('Error loading initial data');
        this.isLoading = false;
      }
    });
  }

  onEventSelect(eventId: number): void {
    this.selectedEventId = eventId;
    if (eventId) {
      this.isLoading = true;
      this.eventAttendeeService.getEventAttendees(eventId).subscribe({
        next: (attendees: User[]) => {
          this.attendees = attendees;
          this.filterAvailableUsers();
          this.isLoading = false;
        },
        error: (err: any) => {
          this.showError('Error loading attendees');
          this.isLoading = false;
        }
      });
    }
  }

  private filterAvailableUsers(): void {
    const attendeeIds = this.attendees.map(a => a.id);
    this.filteredUsers = this.users.filter(user => !attendeeIds.includes(user.id));
  }

  assignGuest(): void {
    if (this.assignForm.valid) {
      this.isLoading = true;
      const { eventId, userId } = this.assignForm.value;

      const newAttendee = {
        event_id: eventId,
        user_id: userId
      };

      this.eventAttendeeService.registerAttendee(newAttendee).subscribe({
        next: () => {
          this.showSuccess('Guest assigned successfully');
          this.onEventSelect(eventId);
          this.assignForm.reset({ eventId });
        },
        error: (err: any) => {
          this.showError('Error assigning guest');
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
          this.onEventSelect(this.selectedEventId!);
        },
        error: (err: any) => {
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