import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';/* 
import { EventService } from '../../services/event.service'; */
import { EventService } from '../../services/eventService.service';
import { Event } from '../../models/Event';

@Component({
  selector: 'app-create-event',
  templateUrl: './create-event.component.html',
  styleUrls: ['./create-event.component.css'],
  standalone: false
})
export class CreateEventComponent implements OnInit {
  eventForm: FormGroup;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private eventService: EventService,
    private router: Router,
    private snackBar: MatSnackBar
  ) {
    this.eventForm = this.fb.group({
      name: ['', [Validators.required, Validators.minLength(3)]],
      description: ['', [Validators.required, Validators.minLength(10)]],
      start_time: ['', [Validators.required, this.dateNotInPast()]],
      end_time: ['', [Validators.required]]
    }, { validators: this.dateRangeValidator });
  }

  private dateNotInPast() {
    return (control: any) => {
      if (control.value) {
        const selectedDate = new Date(control.value);
        const now = new Date();
        if (selectedDate < now) {
          return { dateInPast: true };
        }
      }
      return null;
    };
  }

  private dateRangeValidator(group: FormGroup) {
    const start = group.get('start_time')?.value;
    const end = group.get('end_time')?.value;
    
    if (start && end) {
      const startDate = new Date(start);
      const endDate = new Date(end);
      
      if (endDate <= startDate) {
        return { endBeforeStart: true };
      }
    }
    return null;
  }
  

  ngOnInit(): void {
    const token = localStorage.getItem('Authorization');
    console.log('Token:', token);
    if (!token) {
      console.log('No authorization token found. Please login again.');
      this.router.navigate(['/login']);
      return;
    }
    
    // Verify token format
    if (!token.startsWith('Bearer ')) {
      const formattedToken = `Bearer ${token}`;
      localStorage.setItem('Authorization', formattedToken);
    }
  }

  private getCurrentUserId(): number {
    const token = localStorage.getItem('Authorization');
    if (token) {
      const tokenPayload = JSON.parse(atob(token.split('.')[1]));
      return tokenPayload.user_id;
    }
    return 0;
  }

  onSubmit(): void {
    if (this.eventForm.valid) {
      this.isLoading = true;
      const userId = this.getCurrentUserId();
      
      // Format the dates correctly
      const startTime = new Date(this.eventForm.value.start_time);
      const endTime = new Date(this.eventForm.value.end_time);
      
      const eventData: Omit<Event, 'id' | 'created_at'> = {
        name: this.eventForm.value.name,
        description: this.eventForm.value.description,
        start_time: startTime,
        end_time: endTime,
        created_by: userId
      };

      console.log('Sending event data:', eventData); // Debug log

      this.eventService.createEvent(eventData).subscribe({
        next: (response) => {
          console.log('Response from server:', response); // Debug log
          console.log('Event created successfully');
          this.router.navigate(['/organizer/my-events']);
        },
        error: (error) => {
          console.error('Error details:', error); // Debug log
          if (error.status === 401) {
            console.log('Session expired or unauthorized. Please login again.');
            this.router.navigate(['/login']);
          } else {
            console.log(`Error creating event: ${error.error?.message || error.statusText || 'Unknown error'}`);
          }
          this.isLoading = false;
        },
        complete: () => {
          this.isLoading = false;
        }
      });
    } else {
      console.log('Please fill all required fields correctly');
    }
  }
}
