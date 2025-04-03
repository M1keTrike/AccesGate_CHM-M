import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EventSecurityListComponent } from './event-security-list.component';

describe('EventSecurityListComponent', () => {
  let component: EventSecurityListComponent;
  let fixture: ComponentFixture<EventSecurityListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [EventSecurityListComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EventSecurityListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
