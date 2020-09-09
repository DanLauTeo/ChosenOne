import { Component, OnInit } from '@angular/core';
import { Chatroom } from "../../_models/chatroom";
import { ChatroomService } from '../../_services/chatroom.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-chatlist',
  templateUrl: './chatlist.component.html',
  styleUrls: ['./chatlist.component.css']
})
export class ChatlistComponent implements OnInit {
  chatrooms: Chatroom[];

  constructor(private router: Router, private chatroomService: ChatroomService) { }

  ngOnInit(): void {
    this.getChatrooms();
  }

  getChatrooms(): void {
    this.chatroomService.getChatrooms().subscribe(
      chatrooms => {
        this.chatrooms = chatrooms;
        console.log(this.chatrooms)
      }
    );
  }

  openChat(id): void {
    this.router.navigate(['/chats/'+id]);
  }

}
