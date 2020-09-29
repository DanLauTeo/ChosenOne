import { Component, Inject } from "@angular/core";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { Message } from "../../../_models/message";

@Component({
  selector: "app-delete-message-dialog",
  templateUrl: "./delete-message-dialog.component.html",
  styleUrls: ["./delete-message-dialog.component.css"]
})
export class DeleteMessageDialogComponent {

  constructor(
    public dialogRef: MatDialogRef<DeleteMessageDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public message: Message
  ) {}

  onCancel(): void {
    this.dialogRef.close();
  }

}
