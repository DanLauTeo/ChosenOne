import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FeedImageComponent } from './feed-image.component';

describe('FeedImageComponent', () => {
  let component: FeedImageComponent;
  let fixture: ComponentFixture<FeedImageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FeedImageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FeedImageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
