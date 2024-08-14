# Changelog

## 0.6.0 (2024-08-14)


### ⚠ BREAKING CHANGES

* **minotaur:** ServerActor and ConnActor now use a more streamlined typing system that may affect consumers of the old API.
* **vivid:** ActorId generation and parsing logic has been modified to support cluster identifiers. This may affect existing systems that rely on the previous format.
* **vivid:** GetActorIdByActorRef function has been removed. Update yourcode to use the new Id() method from the ActorRef interface.
* **chrono:** This modifies the internal scheduler logic, which might affect existing clients relying on the previous behavior.
* **vivid:** Mailbox now requires an additional parameter for Enqueue method to specify if the message should be delivered instantly. This may affect the clients relying on the previous signature of the Enqueue method.

### refactor

* **chrono:** update scheduler and task management logic ([4732b99](https://github.com/kercylan98/minotaur/commit/4732b9972719bd2ce9b62715eafa63c199e0d1d8))
* **minotaur:** simplify actor typing and improve network handling ([d542b36](https://github.com/kercylan98/minotaur/commit/d542b3669f3cfb4b113fa027925fbc3efd1f398f))
* **vivid:** optimize actor reference handling and mod status management ([1220b60](https://github.com/kercylan98/minotaur/commit/1220b601bc1675f09db5d0f7fa2de15f0426d4a2))
* **vivid:** optimize message dispatching for instant delivery ([cf23e79](https://github.com/kercylan98/minotaur/commit/cf23e7926adabd23bb6a0158259a717addd4cb5e))


### feat

* **vivid:** add cluster support and refactor actor system ([7d6abf8](https://github.com/kercylan98/minotaur/commit/7d6abf8ccfd8ced1713832bd912706a92329b643))
