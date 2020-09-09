import { Component, OnInit } from '@angular/core';
import { Message } from "../../_models/message";
import { User } from "../../_models/user";
import { ChatroomService } from '../../_services/chatroom.service';
import { MatDividerModule } from '@angular/material/divider';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { AccountService } from '../../_services/account.service'

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})
export class MessagesComponent implements OnInit {
  user : User;
  chatroomName: string;
  id: string;
  messages: Message[];

  constructor(private route: ActivatedRoute, private chatroomService: ChatroomService, private accountService : AccountService) { }

  ngOnInit(): void {
    this.user = this.accountService.getUser();
    this.route.params.subscribe(params => {
      this.id = params['id'];
    });
    this.getMessages();
  }

  getMessages(): void {
    this.chatroomService.getMessages(this.id).subscribe(
      messages => {
        this.messages = messages;
        console.log(this.messages);
      }
    );
  }

  sendMessage(message): void {
    console.log(message);
    this.chatroomService.sendMessage(this.id, message).subscribe(
      messages => {
        this.messages = messages;
        console.log(this.messages);
      }
    );
  }

}
