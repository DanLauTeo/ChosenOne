import { Component, OnInit } from '@angular/core';
import { Image } from '../_models/image';
import { ImageService } from '../_services/image.service'

@Component({
  selector: 'app-feed',
  templateUrl: './feed.component.html',
  styleUrls: ['./feed.component.css']
})
export class FeedComponent implements OnInit {

  images: Image[];

  constructor(private imageService : ImageService) { }

  ngOnInit(): void {
    this.getImages();
  }

  getImages(): void {
    this.imageService.getFeed().subscribe((images) => {
      this.images = images;
      console.log(images)
    });
  }

}
