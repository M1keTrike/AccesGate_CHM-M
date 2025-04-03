import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventAttendeesControlComponent } from './event-attendees-control.component';

describe('EventAttendeesControlComponent', () => {
  let component: EventAttendeesControlComponent;
  let fixture: ComponentFixture<EventAttendeesControlComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [EventAttendeesControlComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EventAttendeesControlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
