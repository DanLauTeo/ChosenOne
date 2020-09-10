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
    
  }

}
