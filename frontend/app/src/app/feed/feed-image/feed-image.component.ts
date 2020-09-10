import { Component, Input, OnInit } from '@angular/core';
import { Image } from '../../_models/image';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-feed-image',
  templateUrl: './feed-image.component.html',
  styleUrls: ['./feed-image.component.css']
})
export class FeedImageComponent implements OnInit {

  @Input() image: Image;

  constructor(private router: Router) { }

  ngOnInit(): void {
  }

  redirect(id) {
    this.router.navigate(['/profile/'+id]);
  }

}
