import { Component, OnInit } from '@angular/core';
import { Chatroom } from "../../_models/chatroom";
import { ChatroomService } from '../../_services/chatroom.service';
import { Router } from '@angular/router';
import { AccountService } from 'src/app/_services/account.service';
import { User } from 'src/app/_models/user';
import { Message } from 'src/app/_models/message';
import { ProfileService } from 'src/app/_services/profile.service';


@Component({
  selector: 'app-chatlist',
  templateUrl: './chatlist.component.html',
  styleUrls: ['./chatlist.component.css']
})
export class ChatlistComponent implements OnInit {

  user: User;
  chatrooms: Chatroom[] = [];
  latestMessages: Message[] = [];
  chatroomNames: string[] = [];

  constructor(
    private router: Router,
    private accountService: AccountService,
    private chatroomService: ChatroomService,
    private profileService: ProfileService,
  ) { }

  ngOnInit(): void {
    this.accountService.user.subscribe(user => {
      this.user = user;
      this.getChatrooms();
    });
  }

  getChatrooms(): void {
    this.chatroomService.getChatrooms().subscribe(
      chatrooms => {
        this.chatrooms = chatrooms;
        this.getLatestMessages();
        this.getChatroomNames();
      }
    );
  }

  getLatestMessages(): void {
    this.latestMessages = [];

    this.chatrooms.forEach((chatroom, idx) => {
      this.latestMessages.push(null);
      if (chatroom.messages && chatroom.messages.length > 0) {
        this.chatroomService.getMessage(chatroom.messages[chatroom.messages.length - 1])
          .subscribe(message => this.latestMessages[idx] = message);
      }
    });
  }

  getChatroomNames(): void {
    this.chatroomNames = [];

    this.chatrooms.forEach((chatroom, idx) => {
      this.chatroomNames.push("");
      let otherUserId = chatroom.participants[0] == this.user.id ? chatroom.participants[1] : chatroom.participants[0];
      this.profileService.getUser(otherUserId)
        .subscribe(user => this.chatroomNames[idx] = user.username);
    });
  }

  openChat(id): void {
    this.router.navigate(['/chats/'+id]);
  }

}
