package  de.alexkrieg.productapp;

import org.springframework.stereotype.Component;
import java.util.List;
import java.util.Arrays;

@Component
class ProductService {
   public List<String> getProducts() {
      return Arrays.asList("iPad","iPod","iPhone");
   }
}