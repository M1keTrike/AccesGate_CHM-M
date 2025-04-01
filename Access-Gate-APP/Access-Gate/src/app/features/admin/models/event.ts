import { User } from './IUsers';

export interface Event {
  id: number;
  name: string;
  description: string;
  start_time: Date;
  end_time: Date;
  created_by: User['id'];
  created_at: Date;
}