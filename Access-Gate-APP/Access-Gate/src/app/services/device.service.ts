import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { AuthService } from './auth.service';

interface DeviceRegistration {
  hardware_id: string;
  assigned_to: number;
  type: string;
  status: string;
  location: string;
}

@Injectable({
  providedIn: 'root'
})
export class DeviceService {
  private apiUrl = `${environment.apiBaseUrl}/devices`;

  constructor(
    private http: HttpClient,
    private authService: AuthService
  ) {}

  private getHeaders(): HttpHeaders {
    const token = this.authService.getToken();
    return new HttpHeaders().set('Authorization', `Bearer ${token}`);
  }

  async verifyDeviceAssignment(mac: string, userId: number): Promise<boolean> {
    try {
      const devices = await this.http.get<any[]>(this.apiUrl, {
        headers: this.getHeaders()
      }).toPromise();
      
      const macExists = devices?.some(device => device.hardware_id === mac) || false;
      if (macExists) {
        alert('Este dispositivo ya está registrado en el sistema'); // Agrega esta líne
        throw new Error('Este dispositivo ya está registrado en el sistema');
      }

      const userHasDevice = devices?.some(device => device.assigned_to === userId) || false;
      if (userHasDevice) {
        throw new Error('Este usuario ya tiene un dispositivo asignado');
      }

      return true;
    } catch (error) {
      throw error;
    }
  }

  async registerDevice(mac: string): Promise<any> {
    const userId = this.authService.getUserId();
    if (!userId) {
      throw new Error('Usuario no autenticado');
    }

    const deviceData: DeviceRegistration = {
      hardware_id: mac,
      assigned_to: userId,
      type: 'ESP32',
      status: 'active',
      location: 'Default'
    };

    return this.http.post(this.apiUrl, deviceData, {
      headers: this.getHeaders()
    }).toPromise();
  }
}