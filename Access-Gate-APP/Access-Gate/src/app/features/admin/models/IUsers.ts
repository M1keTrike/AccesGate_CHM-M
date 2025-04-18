export interface User {
  id: number;
  name: string;
  email: string;
  password_hash: string;
  role: string;
  created_at: string;
  created_by?: number; 
}
