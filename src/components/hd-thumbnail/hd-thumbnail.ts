import { Component, Input, OnInit } from '@angular/core';
import { File } from '../../domain/file';
import { PreviewProvider } from '../../providers/preview/preview';

@Component({
  selector: 'hd-thumbnail',
  templateUrl: 'hd-thumbnail.html'
})
export class HdThumbnailComponent implements OnInit {
  @Input() item: any
  @Input() file: File
  src: string
  constructor(private preview: PreviewProvider) {
    this.src = this.preview.srcDefault
  }
  ngOnInit() {
    if (this.file) {
      this.src = this.preview.getSrc(this.file)
    } else {
      this.preview.itemChanged.subscribe(item => {
        if (item === this.item) {
          this.src = this.preview.getSrc(this.item.f)
        }
      })
      if (this.item) {
        this.src = this.preview.getSrc(this.item.f)
      }
    }
  }
}
