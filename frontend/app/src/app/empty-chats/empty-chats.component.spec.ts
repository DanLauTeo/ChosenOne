import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmptyChatsComponent } from './empty-chats.component';

describe('EmptyChatsComponent', () => {
  let component: EmptyChatsComponent;
  let fixture: ComponentFixture<EmptyChatsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmptyChatsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmptyChatsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
