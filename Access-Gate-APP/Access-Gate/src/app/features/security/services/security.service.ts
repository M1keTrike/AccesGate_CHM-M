import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Event, EventAttendee } from '../models/event-attendee.model';
import { environment } from '../../../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SecurityService {
  private apiUrl = `${environment.apiBaseUrl}`;

  constructor(private http: HttpClient) { }

  private getHeaders() {
    const token = localStorage.getItem('Authorization');
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': `${token}`
      })
    };
  }

  // Obtener todos los eventos
  getAllEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.apiUrl}/events`, this.getHeaders());
  }

  // Obtener asistentes de un evento específico
  getEventAttendees(eventId: number): Observable<EventAttendee[]> {
    return this.http.get<EventAttendee[]>(`${this.apiUrl}/event-attendees/events/${eventId}/attendees`, this.getHeaders());
  }

  // Actualizar el estado de asistencia
  updateAttendanceStatus(eventId: number, userId: number, attended: boolean): Observable<any> {
    return this.http.patch(`${this.apiUrl}/event-attendees/attendance`, {
      event_id: eventId,
      user_id: userId,
      attended: attended
    }, this.getHeaders());
  }
}