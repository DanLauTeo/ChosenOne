import { Component, OnInit, Input } from "@angular/core"
import { ProfileService } from "src/app/_services/profile.service";
import { User } from "src/app/_models/user";

@Component({
  selector: "app-match",
  templateUrl: "./match.component.html",
  styleUrls: ["./match.component.css"],
})
export class MatchComponent implements OnInit {

  private profileService: ProfileService;

  @Input() user_id: string;

  user: User;

  constructor(
    profileService: ProfileService
  ) {
    this.profileService = profileService;
  }

  ngOnInit(): void {
    this.profileService.getUser(this.user_id)
      .subscribe(user => this.user = user)
  }

  startChat(): void {
    console.log(`Start chat with user ${this.user_id}`)
  }
}
