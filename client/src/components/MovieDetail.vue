<template>
    <vue-plyr>
        <video poster="poster.png" src="video.mp4">
            <source src="https://video.fsgn2-2.fna.fbcdn.net/v/t42.9040-2/10000000_1957790207853461_7149694678827270144_n.mp4?_nc_cat=0&efg=eyJybHIiOjE1MDAsInJsYSI6NDA5NiwidmVuY29kZV90YWciOiJzdmVfaGQifQ%3D%3D&rl=1500&vabr=877&oh=f3b581f941ba673aeb965dfb866e3ddd&oe=5B8A89F3"
                    type="video/mp4" size="720">
            <!--<track kind="captions" label="English" srclang="en" src="captions-en.vtt" default>-->
        </video>
    </vue-plyr>
</template>

<script>
    import { phimMoiProvider as movieProvider } from './../services';

    export default {
        name        : 'MovieDetail',
        beforeCreate: function () {
            movieProvider.getMovieByUrl('http://www.phimmoi.net/phim/phien-ngoai-lao-cuu-mon-nhi-nguyet-khai-hoa-7286/xem-phim.html')
        },
    }
</script>

<style scoped>

</style>
