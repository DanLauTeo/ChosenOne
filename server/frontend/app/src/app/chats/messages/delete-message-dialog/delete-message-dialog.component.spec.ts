import { async, ComponentFixture, TestBed } from "@angular/core/testing";

import { DeleteMessageDialogComponent } from "./delete-message-dialog.component";

describe("DeleteMessageDialogComponent", () => {
  let component: DeleteMessageDialogComponent;
  let fixture: ComponentFixture<DeleteMessageDialogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeleteMessageDialogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeleteMessageDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("should create", () => {
    expect(component).toBeTruthy();
  });
});
