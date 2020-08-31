import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GetMatchedComponent } from './get-matched.component';

describe('GetMatchedComponent', () => {
  let component: GetMatchedComponent;
  let fixture: ComponentFixture<GetMatchedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GetMatchedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GetMatchedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
