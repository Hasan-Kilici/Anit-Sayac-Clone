<script setup lang="ts">
import { content } from '#tailwind-config';

const items = [{
  label: 'Giriş',
  content: `
    ‘Anıt Sayaç’, Türkiye’de kadına yönelik şiddetten ölen kadınların anısını yaşatmak için internet üzerinden kurulmuş bir anıt ve her gün güncellenen bir sayaçtır.
    Kadına yönelik şiddet olaylarında 2012 yılının Ocak ve Eylül ayları arasında 125 kadının hayatını kaybettiği açıklandı. Anıt Sayaç kadın cinayetlerinin artarak devam ettiği böyle bir ortamda farkındalık yaratmak ve bilinmeyen verileri açığa kavuşturmak için düşünüldü. Ölen kadınlarımızın isimleriyle anılacağı bu web sitesi, kadına karşı şiddet konusunda toplumun duyarlılığını geliştirme projesi olmanının ötesinde ölen kadınlara adanmış bir anıttır.
    Artış tehdidi tabiatında gizli bu sayaç, şiddetin sürekliliğinin de habercisidir. Kaygı veren bir artış, ağırlaşan bir birikim yanında, aciliyete davet eden bir geri sayım da var ‘Anıt Sayaç’ta. Sayaç attıkça umut eksilmekte; tane tane tükenmektedir. 
    `
}, {
  label: 'Anasayfa',
  content: `‘Anıt Sayaç’ın’ ana sayfası anıtın kendisidir. 2013 yılında kadına yönelik şiddetten öldürülen kadınları gösteren sayaç ve öldürülen kadınların isimlerinden oluşan bir sayfa içeriği düzenlenmiştir. 2008-2015 yılları arasında öldürülen kadınların isimleriyle adeta bir duvar ören bu sayfa hergün güncellenecektir. İsimlerden oluşan bu duvarda, her yıl için değişik bir renk kullanılarak yıllar arasındaki farklılıkları göstermek amaçlanmıştır. Sayfada yer alan her bir kadın ismi aynı zamanda üzerine tıklandığında yeni bir sayfaya açılmaktadır.`
}, {
  label: 'Veriler',
  content: `Bu sitenin veri tabanın oluşturulmasında, 2012’de kadına yönelik şiddetten hayatını kaybeden kadınlara ait veriler konusunda Adalet Bakanlığı, Aile Sosyal Politikalar Bakanlığı ve Emniyet Genel Müdürlüğü’ne yapılan başvuruların sonuçsuz kalması üzerine 2008 – 2013 yılları arasındaki 5 yıllık süre için medya üzerinden yapılan bir tarama yapıldı. Sitede, bu taramanın sonuçlarının “Kadın Cinayetlerini Durduracağız Platformu” ile karşılaştırılmasından sonra ortaya çıkan cinayet listesi kullanıldı. Kadın cinayetleri verileri derlenirken, aile içi şiddet dışı veya kadına yönelik olmayan cinayetler listeye alınmadı. Kadın cinayeti olduğu sonradan anlaşılan cinayetler sayfaya zaman içinde güncellenerek ekleniyor. Tespit edilemeyen cinayetler ‘Şüpheli Ölüm’ olarak listede yer alıyor. Anıtsayaç sitesinde sadece medya da çıkan cinayetler kullanılıyor. Medya’nın bilmediği, bizim ulaşamadığımız bir sürü cinayetin olması mümkündür.

Ölen her bir kadın için ayrı bir sayfanın açılması ile birlikte kadının neden, kim tarafından öldürüldüğü, nasıl öldürüldüğü, devletten koruma talebinde bulunup bulunmadığı gibi bilgilerin yanı sıra, cinayet hakkında medyada çıkan haberlere de ulaşılabilmektedir. Böylelikle her geçen yıl medyanın olayları ele alışındaki dil farkına da dikkat çekilmektedir.

anitsayac.com 2013 yılından bu yana ‘Kadın Cinayetlerini Durduracağız Platformu’ ile birlikte, verilerin oluşturulması kapsamında işbirliğine devam ediyor.`
},{
    label:"Arama",
    content: `Ana sayfada kadına yönelik şiddet yüzünden hayatını kaybeden belirli bir kadını bulmanız için oluşturuldu. Aramak istediğiniz ismi yazarak o kişiyle ilgili web sayfasına ulaşabilirsiniz.`
},{
    label: "İletişim",
    content: "iletisim@anitsayac.com",
}]
</script>

<template>
    <div style="width: max(100% - 40px);">
        <UTabs :items="items" orientation="vertical" :ui="{ 
            wrapper: 'flex items-center gap-4 max-h-[100dvh] h-[100dvh] items-stretch py-4 px-4', 
            list: { width: 'w-100', background: 'bg-transparent', }, 
        }" />
    </div>
</template>