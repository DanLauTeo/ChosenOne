import { Component, OnInit, Input, OnChanges } from "@angular/core"
import { ProfileService } from "../../_services/profile.service";
import { User } from "../../_models/user";

@Component({
  selector: "app-match",
  templateUrl: "./match.component.html",
  styleUrls: ["./match.component.css"],
})
export class MatchComponent implements OnInit {

  private profileService: ProfileService;

  @Input() user_id: string;

  user: User;

  constructor(profileService: ProfileService) {
    this.profileService = profileService;
  }

  ngOnInit(): void {
  }

  ngOnChanges(): void {
    this.profileService.getUser(this.user_id)
      .subscribe(user => this.user = user);
  }

  openProfile(): void {
    console.log(`Open profile of ${this.user.id}`);
  }

  startChat(event: Event): void {
    console.log(`Start chat with user ${this.user.id}`);
    event.stopPropagation();
  }
}