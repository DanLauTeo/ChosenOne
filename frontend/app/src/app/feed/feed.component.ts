import { Component, OnInit } from '@angular/core';
import { Image } from '../_models/image';

@Component({
  selector: 'app-feed',
  templateUrl: './feed.component.html',
  styleUrls: ['./feed.component.css']
})
export class FeedComponent implements OnInit {

  images: Image[];

  constructor() { }

  ngOnInit(): void {
    this.images = [{imgURL:"https://images.unsplash.com/photo-1494548162494-384bba4ab999?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80", ownerID:"1", ownerName:"John Smith", ownerProfile:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQP-i5liksKo3g85Qz90jpYieJ4J_YGy5S7JQ&usqp=CAU"},
      {imgURL:"https://image.shutterstock.com/image-photo/mountains-during-sunset-beautiful-natural-260nw-407021107.jpg", ownerID:"2", ownerName:"Joe Bloggs", ownerProfile:"https://www.jumpstarttech.com/files/2018/08/Network-Profile.png"},
      {imgURL:"https://cdn.jpegmini.com/user/images/slider_puffin_jpegmini_mobile.jpg", ownerID:"3", ownerName:"Mike Brown", ownerProfile:"https://meetanentrepreneur.lu/wp-content/uploads/2019/08/profil-linkedin-300x300.jpg"}
    ]
  }

}
