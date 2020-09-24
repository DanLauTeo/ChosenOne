import { Component, OnInit } from '@angular/core';
import { Message } from "../../_models/message";
import { User } from "../../_models/user";
import { ChatroomService } from '../../_services/chatroom.service';
import { MatDividerModule } from '@angular/material/divider';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { AccountService } from '../../_services/account.service'
import { ProfileService } from 'src/app/_services/profile.service';
import { Observable } from 'rxjs';
import { map, flatMap } from 'rxjs/operators';
import { Chatroom } from 'src/app/_models/chatroom';

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})
export class MessagesComponent implements OnInit {
  user: User;
  otherUser: User;
  chatroom: Chatroom;
  chatroomId: number;
  messages: Message[];

  constructor(
    private route: ActivatedRoute,
    private chatroomService: ChatroomService,
    private accountService: AccountService,
    private profileService: ProfileService,
  ) {
    this.user = this.accountService.getUser();
  }

  ngOnInit(): void {
    this.route.params.subscribe(params => {
      this.chatroomId = +params['id'];
      this.getUserInfo();
    });
  }

  getUserInfo(): void {
    this.accountService.user
      .subscribe(user => {
        this.user = user;
        this.getChatroomInfo();
      });
  }

  getChatroomInfo(): void {
    this.chatroomService.getChatroom(this.chatroomId)
      .subscribe(chatroom => {
          this.chatroom = chatroom;
          this.getOtherUserInfo();
      });
  }

  getOtherUserInfo(): void {
    let userId = this.user.id;
    let otherUserId = this.chatroom.participants.filter(x => x != userId)[0];
    this.profileService.getUser(otherUserId)
      .subscribe(otherUser => {
        this.otherUser = otherUser
        this.getMessages();
      });
  }

  getMessages(): void {
    this.chatroomService.getMessages(this.chatroomId).subscribe(
      messages => {
        this.messages = messages;
      }
    );
  }

  sendMessage(message): void {
    this.chatroomService.sendMessage(this.chatroomId, message).subscribe(
      _ => this.getMessages()
    );
  }

  getSenderUsername(message: Message) {
    return message.sender_id == this.user.id ? this.user.username : this.otherUser.username;
  }

  deleteMessage(id: number): void {
    this.chatroomService.deleteMessage(id)
      .subscribe(_ => this.getMessages());
  }

}
