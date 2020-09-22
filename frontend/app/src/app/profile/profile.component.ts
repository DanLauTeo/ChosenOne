import { Component, OnInit } from '@angular/core';
import { User } from '../_models/user';
import { Patch } from '../_models/patch';
import { AccountService } from '../_services/account.service'
import { ProfileService } from '../_services/profile.service'
import { ActivatedRoute, Router } from '@angular/router';
import { ChatroomService } from '../_services/chatroom.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  user : User;
  isCurrentUser: boolean;
  isEditable: boolean;
  id: string;
  loggedIn: boolean;

  constructor( private accountService : AccountService, private chatroomService : ChatroomService, private profileService : ProfileService, private route : ActivatedRoute, private router: Router) {

  }

  ngOnInit(): void {
    this.isCurrentUser = false;
    this.isEditable = false;
    this.loggedIn = false;
    this.getProfile();
  }

  getProfile(): void {
    //Get ID from route
    this.id = this.route.snapshot.paramMap.get('id');

    //Check if user logged in
    if(this.accountService.getUser()==null){
      //if true then user not logged in, can still view profile if route has ID
      if (this.id == null) {
        this.router.navigate(['/login'])
      }
    } else {
      this.loggedIn = true;
      //Get ID of logged in user
      let currentId = this.accountService.getUserID();

      //If no ID in route, user is on own profile
      if (this.id == null) {
        this.id = currentId;
      }

      //If user on own profile, set isCurrentUser to true
      if (currentId == this.id) {
        this.isCurrentUser = true;
      }
    }

    this.profileService.getUser(this.id).subscribe((user) => {
      this.user = user;
    });
  }

  patchProfile(): void {
    //Get ID from route
    let id = this.route.snapshot.paramMap.get('id');

    //Get ID of logged in user
    let currentId = this.accountService.getUserID();

    //If no ID in route, user is on own profile
    if (id == null) {
      id = currentId;
    }

    //If user not on own profile, return without changing
    if (!this.isCurrentUser) {
      console.log("Not on own profile");
      return;
    }
    let opArray: Patch[] = [];
    let newName = document.getElementById("username").innerHTML;
    let newBio = document.getElementById("bio").innerHTML;
    //Get new username
    if (this.user.username != newName) {
      opArray.push({op:"replace", path:"/Name", value: newName});
    }
    if (this.user.bio != newBio) {
      opArray.push({op:"replace", path:"/Bio", value: newBio});
    }
    this.profileService.patchUser(this.id, opArray).subscribe((user) => {
      this.user = user;
    });
  }

  edit() {
    if (this.isCurrentUser){
      if(document.getElementById("username").contentEditable ==  "true"){
        this.patchProfile()
        document.getElementById("username").contentEditable = "false";
        document.getElementById("bio").contentEditable = "false";
        this.isEditable = false;
      } else {
        this.isEditable = true;
        document.getElementById("username").contentEditable = "true";
        document.getElementById("bio").contentEditable = "true";
      }
    }
  }

  onPicChange(event) {
    this.user = event;
  }

  logout() {
    this.accountService.didLogout();
    this.router.navigate(["/login"]);
  }

  openChat(event: Event): void {
    this.chatroomService.startChat(this.id)
      .subscribe(chatroom => this.router.navigate([`/chats/${chatroom.id}/`]));
    event.stopPropagation();
  }
}
