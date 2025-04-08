import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { SecurityService } from '../../services/security.service';
import { SecurityWebSocketService } from '../../services/security-websocket.service';
import { EventAttendee } from '../../models/event-attendee.model';

@Component({
  selector: 'app-event-attendees-control',
  templateUrl: './event-attendees-control.component.html',
  styleUrls: ['./event-attendees-control.component.css'],
  standalone: false
})
export class EventAttendeesControlComponent implements OnInit, OnDestroy {
  eventId: number | null = null;
  attendees: EventAttendee[] = [];
  loading: boolean = true;
  lastScannedNFC: string | null = null;

  constructor(
    private route: ActivatedRoute,
    private securityService: SecurityService,
    private wsService: SecurityWebSocketService
  ) {}

  ngOnInit(): void {
    this.eventId = Number(this.route.snapshot.paramMap.get('id'));
    if (this.eventId) {
      this.loadAttendees();
      this.initializeWebSocket();
    }
  }

  ngOnDestroy() {
    this.wsService.disconnect();
  }

  private initializeWebSocket() {
    const deviceMac = localStorage.getItem('device_mac');
    if (!deviceMac) {
      console.error('❌ No se encontró la dirección MAC del dispositivo');
      return;
    }

    this.wsService.connect('nfc', deviceMac);
    this.wsService.getMessages().subscribe({
      next: (data) => {
        this.handleNFCScan(data.nfcId);
      },
      error: (error) => {
        console.error('Error en WebSocket:', error);
      }
    });
  }

  private handleNFCScan(nfcId: string) {
    this.lastScannedNFC = nfcId;
    console.log('📱 NFC escaneado:', nfcId);

    // Buscar el asistente que coincida con el NFC escaneado
    const attendee = this.attendees.find(a => a.nfc_id === nfcId);
    if (attendee && !attendee.attended) {
      this.updateAttendanceStatus(attendee.user_id, true);
    }
  }

  loadAttendees(): void {
    if (!this.eventId) return;
    
    this.securityService.getEventAttendees(this.eventId).subscribe({
      next: (attendees) => {
        this.attendees = attendees;
        this.loading = false;
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
        console.log('✅ Estado de asistencia actualizado');
        this.loadAttendees();
      },
      error: (error) => {
        console.error('❌ Error actualizando estado de asistencia:', error);
      }
    });
  }
}
