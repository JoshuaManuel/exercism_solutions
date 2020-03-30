#include "leap.h"

/*
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
    unless the year is also evenly divisible by 400
*/


namespace leap {
    bool is_leap_year(int y) {
        if (y%4 == 0) {
            if (y%100 == 0) {
                if (y%400 == 0) {
                    return true;
                }
                return false;
            }
            return true;
        }
        return false;
    };
}  // namespace leap
