import { Component, Input, OnInit } from '@angular/core';
import { Image } from '../../_models/image';

@Component({
  selector: 'app-feed-image',
  templateUrl: './feed-image.component.html',
  styleUrls: ['./feed-image.component.css']
})
export class FeedImageComponent implements OnInit {

  @Input() image: Image;

  constructor() { }

  ngOnInit(): void {
  }

  redirect(id) {
    console.log("Click!")
  }

}
