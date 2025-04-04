import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SecurityWebSocketService {
  private socket!: WebSocket;
  private messageSubject = new Subject<any>();
  private baseUrl = 'wss://access-gate-ws.acstree.xyz/ws';

  constructor() {}

  connect(topic: string, mac?: string) {
    let wsUrl = `${this.baseUrl}?topic=${topic}`;
    if (mac) {
      wsUrl += `&mac=${mac}`;
    }

    if (this.socket) {
      this.socket.close();
    }

    this.socket = new WebSocket(wsUrl);

    this.socket.onopen = () => {
      console.log(`✅ Conectado al WebSocket de Seguridad - Topic: ${topic}`);
    };

    this.socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        const content = JSON.parse(data.content);
        this.messageSubject.next({
          nfcId: content.evento,
          mac: content.mac
        });
      } catch (error) {
        console.error('❌ Error al procesar el mensaje:', error);
      }
    };

    this.socket.onerror = (error) => {
      console.error('❌ Error en WebSocket:', error);
    };

    this.socket.onclose = () => {
      console.log('🔌 Conexión WebSocket cerrada');
    };
  }

  disconnect() {
    if (this.socket) {
      this.socket.close();
    }
  }

  getMessages(): Observable<any> {
    return this.messageSubject.asObservable();
  }

  isConnected(): boolean {
    return this.socket && this.socket.readyState === WebSocket.OPEN;
  }
}