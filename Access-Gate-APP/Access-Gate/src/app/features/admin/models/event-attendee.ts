import { User } from './IUsers';
import { Event } from './event';

export interface EventAttendee {
  id: number;
  user_id: User['id'];
  event_id: Event['id'];
  registered_at: Date;
}