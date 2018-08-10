import { Component, Input, Output, EventEmitter } from '@angular/core';
import { File } from '../../domain/file'
@Component({
  selector: 'hd-files',
  templateUrl: 'hd-files.html'
})
export class HdFilesComponent {
  @Input() files: File[]
  @Output() selectFile = new EventEmitter<File>()
  constructor() {
  }
  onSelectFile(file: File) {
    this.selectFile.emit(file)
  }
  onDownload(file: File) {
    const a = document.createElement("a")
    a.href = `/down/${file.id}`
    a.download = file.name
    a.click()
    a.remove()
  }
}
