import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { EventService } from '../../services/eventService.service';
import { Event } from '../../models/Event';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';

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
    private fb: FormBuilder,
    private snackBar: MatSnackBar
  ) {
    this.eventForm = this.fb.group({
      name: ['', [Validators.required, Validators.minLength(3)]],
      description: ['', [Validators.required, Validators.minLength(10)]],
      start_time: ['', [Validators.required, this.dateNotInPast()]],
      end_time: ['', [Validators.required]]
    }, { validators: this.dateRangeValidator });
  }

  ngOnInit() {
    this.eventId = Number(this.route.snapshot.paramMap.get('id'));
    this.loadEvent();
    console.log("id del evento: ",this.eventId);
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
      this.showError('Por favor, complete todos los campos correctamente');
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
        this.showSuccess('Evento actualizado exitosamente');
        this.router.navigate(['/organizer/my-events']);
      },
      error: (error) => {
        if (error.status === 401) {
          this.showError('Sesión expirada o no autorizada. Por favor, inicie sesión nuevamente.');
          this.router.navigate(['/login']);
        } else {
          this.showError(`Error al actualizar el evento: ${error.error?.message || error.statusText || 'Error desconocido'}`);
        }
      }
    });
  }

  private showSuccess(message: string): void {
    this.snackBar.open(message, 'Cerrar', { 
      duration: 3000,
      panelClass: ['success-snackbar']
    });
  }

  private showError(message: string): void {
    this.snackBar.open(message, 'Cerrar', { 
      duration: 3000,
      panelClass: ['error-snackbar']
    });
  }
}
