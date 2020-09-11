import { Component, OnInit } from '@angular/core';
import { User } from '../_models/user';
import { Patch } from '../_models/patch';
import { AccountService } from '../_services/account.service'
import { Observable } from 'rxjs';
import { USER } from '../_mock_models/mock_user'
import { ProfileService } from '../_services/profile.service'
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {
  user : User;
  isCurrentUser: boolean;


  constructor( private accountService : AccountService, private profileService : ProfileService, private route : ActivatedRoute) {
    //this.accountService.user.subscribe(x => this.user = x);
    this.user = accountService.getUser();
  }

  ngOnInit(): void {
    this.isCurrentUser = false;
    this.getProfile();
  }

  getProfile(): void {
    //Get ID from route
    let id = this.route.snapshot.paramMap.get('id');

    //Get ID of logged in user
    let currentId = this.accountService.getUserID();

    //If no ID in route, assume user is going to own profile
    if (id == null) {
      id = currentId;
    } 

    //If user on own profile, set isCurrentUser to true
    if (currentId == id) {
      this.isCurrentUser = true;
    }

    this.profileService.getUser(id).subscribe((user) => {
      this.user = user;
    });
  }

  patchProfile(): void {
    console.log("Called");
    //Get ID from route
    let id = this.route.snapshot.paramMap.get('id');

    //Get ID of logged in user
    let currentId = this.accountService.getUserID();

    //If no ID in route, user is on own profile
    if (id == null) {
      id = currentId;
    } 

    //If user not on own profile, return without changing
    if (currentId != id) {
      console.log("ID's dont match");
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
    this.profileService.patchUser(id, opArray).subscribe((user) => {
      this.user = user;
    });
  }
    
  edit() {
    if (this.isCurrentUser){
      if(document.getElementById("username").contentEditable ==  "true"){
        this.patchProfile()
        document.getElementById("username").contentEditable = "false";
        document.getElementById("bio").contentEditable = "false";
      } else {
        document.getElementById("username").contentEditable = "true";
        document.getElementById("bio").contentEditable = "true";
      }
    }
  }

  logout() {
    this.accountService.logout();
  }
}


//i want you to send an message object
//at message/chatroomID
//with(senderID, body, timestamp)
