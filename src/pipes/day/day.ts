import { Pipe, PipeTransform } from '@angular/core';
import dayjs from 'dayjs'

@Pipe({
  name: 'day',
})
export class DayPipe implements PipeTransform {
  /**
   * 日志显示
   */
  transform(value: string, ...args) {
    const now = dayjs()
    if (value === now.format('YYYY-MM-DD')) {
      return '今天'
    }
    if (value === now.add(-1, 'day').format('YYYY-MM-DD')) {
      return '昨天'
    }
    if (value === now.add(-2, 'day').format('YYYY-MM-DD')) {
      return '前天'
    }
    const d = dayjs(value)
    if (d.year() === now.year()) {
      if (d.month() == now.month()) {
        return d.format('D日')
      }
      return d.format('M月D日')
    }
    return d.format('YYYY年M月D日')
  }
}
