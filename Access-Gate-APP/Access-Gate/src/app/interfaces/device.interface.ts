export interface Device {
  hardware_id: string;  // MAC address
  type: string;
  status: string;
  location: string;
  assigned_to: number;
}