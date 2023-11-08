/* import the fontawesome core */
import { library } from "@fortawesome/fontawesome-svg-core";

/* import font awesome icon component */
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

/* import specific icons */
import {
  faPlus,
  faCopy,
  faArrowRight,
  faBars,
  faSignal,
  faXmark,
} from "@fortawesome/free-solid-svg-icons";

/* add icons to the library */
library.add(faPlus, faCopy, faArrowRight, faBars, faSignal, faXmark);

export default FontAwesomeIcon;
