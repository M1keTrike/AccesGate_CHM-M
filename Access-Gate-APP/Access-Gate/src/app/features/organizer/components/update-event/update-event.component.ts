import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EventService } from '../../services/eventService.service';
import { Event } from '../../models/Event';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-update-event',
  templateUrl: './update-event.component.html',
  styleUrls: ['./update-event.component.css'],
  standalone: false
})
export class UpdateEventComponent implements OnInit {
  eventForm: FormGroup;
  eventId: number | null = null;
  loading = true;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private eventService: EventService,
    private fb: FormBuilder
  ) {
    this.eventForm = this.fb.group({
      name: ['', Validators.required],
      description: ['', Validators.required],
      start_time: ['', Validators.required],
      end_time: ['', Validators.required]
    });
  }

  ngOnInit() {
    this.eventId = Number(this.route.snapshot.paramMap.get('id'));
    this.loadEvent();
    console.log("id del evento: ",this.eventId);
  }

  private loadEvent() {
    if (!this.eventId) {
      console.error('Event ID is missing');
      this.loading = false;
      return;
    }
    this.eventService.getEventById(this.eventId).subscribe({
      next: (event) => {
        this.eventForm.patchValue({
          name: event.name,
          description: event.description,
          start_time: event.start_time,
          end_time: event.end_time
        });
        this.loading = false;
      },
      error: (error) => {
        console.error('Error loading event:', error);
        this.loading = false;
      }
    });
  }

  onSubmit() {
    if (!this.eventId || !this.eventForm.valid) {
      console.log('Please fill all required fields correctly');
      return;
    }

    const userId = parseInt(localStorage.getItem('userId') || '1', 10);
    
    // Format the dates correctly
    const startTime = new Date(this.eventForm.value.start_time);
    const endTime = new Date(this.eventForm.value.end_time);

    const updatedEvent: Event = {
      id: this.eventId,
      name: this.eventForm.value.name,
      description: this.eventForm.value.description,
      start_time: startTime,
      end_time: endTime,
      created_by: userId,
      created_at: new Date()
    };

    console.log('Sending updated event data:', updatedEvent);

    this.eventService.updateEvent(this.eventId, updatedEvent).subscribe({
      next: (response) => {
        console.log('Response from server:', response);
        console.log('Event updated successfully');
        this.router.navigate(['/organizer/my-events']);
      },
      error: (error) => {
        console.error('Error details:', error);
        if (error.status === 401) {
          console.log('Session expired or unauthorized. Please login again.');
          this.router.navigate(['/login']);
        } else {
          console.log(`Error updating event: ${error.error?.message || error.statusText || 'Unknown error'}`);
        }
      }
    });
  }
}
