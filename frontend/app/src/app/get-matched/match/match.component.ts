import { Component, OnInit, Input, OnChanges } from "@angular/core"
import { ProfileService } from "../../_services/profile.service";
import { User } from "../../_models/user";
import { ChatroomService } from 'src/app/_services/chatroom.service';
import { Router } from '@angular/router';

@Component({
  selector: "app-match",
  templateUrl: "./match.component.html",
  styleUrls: ["./match.component.css"],
})
export class MatchComponent implements OnInit {

  @Input() userId: string;

  user: User;

  constructor(
    private router: Router,
    private profileService: ProfileService,
    private chatroomService: ChatroomService,
  ) {
    this.profileService = profileService;
  }

  ngOnInit(): void {
  }

  ngOnChanges(): void {
    this.profileService.getUser(this.userId)
      .subscribe(user => this.user = user);
  }

  openProfile(): void {
    console.log(`Open profile of ${this.user.id}`);
  }

  startChat(event: Event): void {
    this.chatroomService.startChat(this.userId)
      .subscribe(chatroom => this.router.navigate([`/chats/${chatroom.id}/`]));
    event.stopPropagation();
  }
}